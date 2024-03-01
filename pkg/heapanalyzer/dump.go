package heapanalyzer

import (
	"fmt"
	"math"
	"sort"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/grafana/pyroscope/pkg/heapanalyzer/debug/core"
	"github.com/grafana/pyroscope/pkg/heapanalyzer/debug/gocore"
)

type Dump struct {
	l        log.Logger
	exePath  string
	corePath string
	core     *core.Process
	gocore   *gocore.Process
	info     *HeapDump
}

func NewDump(l log.Logger, exePath string, corePath string, info *HeapDump) (*Dump, error) {
	c, err := core.Core(corePath, "", exePath)
	if err != nil {
		return nil, err
	}
	p, err := gocore.Core(c)
	if err != nil {
		return nil, err
	}
	d := &Dump{
		l:        l,
		exePath:  exePath,
		corePath: corePath,
		core:     c,
		gocore:   p,
		info:     info,
	}
	err = d.InitHeap()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Dump) InitHeap() (err error) {
	defer func() {
		if r := recover(); r != nil {
			level.Error(d.l).Log("msg", "recovered from panic", "panic", r)
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	d.gocore.TypeHeap()
	return err
}

// ObjectFields returns all fields of a heap object.
func (d *Dump) ObjectFields(obj int64) ([]*Field, error) {
	o, err := d.findObject(obj)
	if err != nil {
		return nil, err
	}

	fields := make([]*Field, 0)

	var end int64
	if o.typ != nil {
		n := o.size / o.typ.Size
		if n > 1 {
			for i := int64(0); i < n; i++ {
				fields = d.getFields(d.gocore, fmt.Sprintf("[%d]", i), o.addr.Add(i*o.typ.Size), o.typ, fields)
			}
		} else {
			fields = d.getFields(d.gocore, "", o.addr, o.typ, fields)
		}
		end = n * o.typ.Size
	}

	for i := end; i < o.size; i += d.gocore.Process().PtrSize() {
		// fmt.Fprintf(w, "<tr><td>f%d</td><td colspan=\"2\">?</td>", i)
		if d.gocore.IsPtr(o.addr.Add(i)) {
			// fmt.Fprintf(w, "<td>%s</td>", htmlPointer(c, c.Process().ReadPtr(addr.Add(i))))
		} else {
			// fmt.Fprintf(w, "<td><pre>")
			for j := int64(0); j < d.gocore.Process().PtrSize(); j++ {
				// fmt.Fprintf(w, "%02x ", c.Process().ReadUint8(addr.Add(i+j)))
			}
			// fmt.Fprintf(w, "</pre></td><td><pre>")
			for j := int64(0); j < d.gocore.Process().PtrSize(); j++ {
				r := d.gocore.Process().ReadUint8(o.addr.Add(i + j))
				if r >= 32 && r <= 126 {
					// fmt.Fprintf(w, "%s", html.EscapeString(string(rune(r))))
				} else {
					// fmt.Fprintf(w, ".")
				}
			}
			// fmt.Fprintf(w, "</pre></td>")
		}
		// fmt.Fprintf(w, "</tr>\n")
	}

	return fields, nil
}

// this is a copy of the pyroscope/pkg/heapanalyzer/debug/cmd/viewcore/html.go:htmlObject
// TODO: rename to more appropriate name
// the idea is to instead of writing to a writer, we append the field to the list of fields
// c is a d.gocore and so in
func (d *Dump) getFields(c *gocore.Process, name string, a core.Address, t *gocore.Type, fields []*Field) []*Field {
	switch t.Kind {
	case gocore.KindBool:
		v := c.Process().ReadUint8(a) != 0
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%t", v),
		})
	case gocore.KindInt:
		var v int64
		switch t.Size {
		case 1:
			v = int64(c.Process().ReadInt8(a))
		case 2:
			v = int64(c.Process().ReadInt16(a))
		case 4:
			v = int64(c.Process().ReadInt32(a))
		case 8:
			v = c.Process().ReadInt64(a)
		}
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%d", v),
		})
	case gocore.KindUint:
		var v uint64
		switch t.Size {
		case 1:
			v = uint64(c.Process().ReadUint8(a))
		case 2:
			v = uint64(c.Process().ReadUint16(a))
		case 4:
			v = uint64(c.Process().ReadUint32(a))
		case 8:
			v = c.Process().ReadUint64(a)
		}
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%d", v),
		})
	case gocore.KindFloat:
		var v float64
		switch t.Size {
		case 4:
			v = float64(math.Float32frombits(c.Process().ReadUint32(a)))
		case 8:
			v = math.Float64frombits(c.Process().ReadUint64(a))
		}
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%f", v),
		})
	case gocore.KindComplex:
		var v complex128
		switch t.Size {
		case 8:
			v = complex128(complex(
				math.Float32frombits(c.Process().ReadUint32(a)),
				math.Float32frombits(c.Process().ReadUint32(a.Add(4)))))

		case 16:
			v = complex(
				math.Float64frombits(c.Process().ReadUint64(a)),
				math.Float64frombits(c.Process().ReadUint64(a.Add(8))))
		}
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%f", v),
		})
	case gocore.KindEface:
		fields = append(fields, &Field{
			Name:    name,
			Type:    "interface{}",
			Value:   fmt.Sprintf("%x", c.Process().ReadPtr(a)),
			Pointer: fmt.Sprintf("unsafe.Pointer [%x]", a.Add(c.Process().PtrSize())),
		})
	case gocore.KindIface:
		dt := c.DynamicType(t, a)
		if dt != nil {
			fields = append(fields, &Field{
				Name:    name,
				Type:    fmt.Sprintf("interface{...} %s", dt.Name),
				Value:   fmt.Sprintf("%x", c.Process().ReadPtr(a)),
				Pointer: fmt.Sprintf("unsafe.Pointer [%x]", a.Add(c.Process().PtrSize())),
			})
		}
	case gocore.KindPtr:
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%x", c.Process().ReadPtr(a)),
		})
	case gocore.KindFunc:
		if fn := c.Process().ReadPtr(a); fn != 0 {
			pc := c.Process().ReadPtr(fn)
			if f := c.FindFunc(pc); f != nil && f.Entry() == pc {
				fields = append(fields, &Field{
					Name:    name,
					Type:    t.String(),
					Value:   f.Name(),
					Pointer: fmt.Sprintf("%x", c.Process().ReadPtr(a)),
				})
			}
		}
	case gocore.KindString:
		n := c.Process().ReadInt(a.Add(c.Process().PtrSize()))
		if n > 0 {
			n2 := n
			ddd := ""
			if n > 100 {
				n2 = 100
				ddd = "..."
			}
			b := make([]byte, n2)
			c.Process().ReadAt(b, c.Process().ReadPtr(a))
			fields = append(fields, &Field{
				Name:    name,
				Type:    "string",
				Value:   fmt.Sprintf("%s%s", string(b), ddd),
				Pointer: fmt.Sprintf("%x", c.Process().ReadPtr(a)),
			})
		} else {
			fields = append(fields, &Field{
				Name:    name,
				Type:    "string",
				Value:   "",
				Pointer: fmt.Sprintf("%x", c.Process().ReadPtr(a)),
			})
		}
	case gocore.KindSlice:
		fields = append(fields, &Field{
			Name:    name,
			Type:    t.String(),
			Value:   fmt.Sprintf("%d", c.Process().ReadInt(a.Add(c.Process().PtrSize()))),
			Pointer: fmt.Sprintf("%d", c.Process().ReadInt(a.Add(c.Process().PtrSize()*2))),
		})
	case gocore.KindArray:
		fields = append(fields, &Field{
			Name:  name,
			Type:  t.String(),
			Value: fmt.Sprintf("%d", t.Count),
		})
	case gocore.KindStruct:
		fields2 := make([]*Field, 0)
		for _, f := range t.Fields {
			fields2 = d.getFields(c, name+"."+f.Name, a.Add(f.Off), f.Type, fields2)
		}
		fields = append(fields, &Field{
			Name:   name,
			Type:   t.String(),
			Value:  fmt.Sprintf("%x", c.Process().ReadPtr(a)),
			Fields: fields2,
		})
	default:
		level.Warn(d.l).Log("msg", "unsupported type", "type", t.Kind.String())
	}
	return fields
}

func (d *Dump) findObject(obj int64) (*object, error) {
	a := core.Address(obj)
	x, _ := d.gocore.FindObject(a)
	if x == 0 {
		return nil, fmt.Errorf("can't find object at %x", a)
	}

	addr := d.gocore.Addr(x)
	size := d.gocore.Size(x)
	typ, repeat := d.gocore.Type(x)

	return &object{
		addr:   addr,
		size:   size,
		typ:    typ,
		repeat: repeat,
	}, nil
}

// Object returns all heap objects.
func (d *Dump) Objects() []*Object {
	var buckets []*Object

	d.gocore.ForEachObject(func(x gocore.Object) bool {
		addr := fmt.Sprintf("%x", d.gocore.Addr(x))
		typeName := typeName(d.gocore, x)

		buckets = append(buckets, &Object{
			Id:          addr, // TODO: use real id
			Type:        typeName,
			Address:     addr,
			DisplayName: typeName + " [" + addr + "]", // TODO: use real display name
			Size:        d.gocore.Size(x),
		})

		return true
	})

	return buckets
}

// ObjectTypes returns a list of object types in the heap, sorted by total size.
func (d *Dump) ObjectTypes() []*ObjectTypeStats {
	level.Debug(d.l).Log("msg", "calculating object types")

	var buckets []*ObjectTypeStats
	m := map[string]*ObjectTypeStats{}

	d.gocore.ForEachObject(func(x gocore.Object) bool {
		name := typeName(d.gocore, x)
		b := m[name]
		if b == nil {
			b = &ObjectTypeStats{Type: name, TotalSize: d.gocore.Size(x), Count: 1}
			buckets = append(buckets, b)
			m[name] = b
		} else {
			b.Count++
			b.TotalSize += d.gocore.Size(x)
		}

		return true
	})

	level.Debug(d.l).Log("msg", "calculated object types", "count", len(buckets))

	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].TotalSize*buckets[i].Count > buckets[j].TotalSize*buckets[j].Count
	})

	return buckets
}

// typeName returns a string representing the type of this object.
func typeName(c *gocore.Process, x gocore.Object) string {
	size := c.Size(x)
	typ, repeat := c.Type(x)
	if typ == nil {
		return fmt.Sprintf("unk%d", size)
	}

	name := typ.String()
	n := size / typ.Size
	if n > 1 {
		if repeat < n {
			name = fmt.Sprintf("[%d+%d?]%s", repeat, n-repeat, name)
		} else {
			name = fmt.Sprintf("[%d]%s", repeat, name)
		}
	}
	return name
}

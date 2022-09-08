// This component is based on logic from the flamebearer project
// https://github.com/mapbox/flamebearer

// ISC License

// Copyright (c) 2018, Mapbox

// Permission to use, copy, modify, and/or distribute this software for any purpose
// with or without fee is hereby granted, provided that the above copyright notice
// and this permission notice appear in all copies.

// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
// OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.
import { css } from '@emotion/css';
import React, { useCallback, useEffect, useMemo, useRef } from 'react';
import { useWindowSize } from 'react-use';

import { DataFrame, DataFrameView } from '@grafana/data';
import { useStyles2 } from '@grafana/ui';

import { COLLAPSE_THRESHOLD, PIXELS_PER_LEVEL } from '../../constants';
import { getBarX, getRectDimensionsForLevel, renderRect } from './rendering';
import { Item, ItemWithStart, nestedSetToLevels } from './dataTransform';

type Props = {
  data: DataFrame;
  topLevelIndex: number;
  rangeMin: number;
  rangeMax: number;
  query: string;
  setTopLevelIndex: (level: number) => void;
  setRangeMin: (range: number) => void;
  setRangeMax: (range: number) => void;
};

const FlameGraph = ({
  data,
  topLevelIndex,
  rangeMin,
  rangeMax,
  query,
  setTopLevelIndex,
  setRangeMin,
  setRangeMax,
}: Props) => {
  const styles = useStyles2(getStyles);

  const totalTicks = data.fields[1].values.get(0);

  const levels = useMemo(() => {
    if (!data) {
      return [];
    }
    const dataView = new DataFrameView<Item>(data);
    return nestedSetToLevels(dataView);
  }, [data]);
  console.log(levels);
  console.log(data);

  const { width: windowWidth } = useWindowSize();
  const graphRef = useRef<HTMLCanvasElement>(null);

  // convert pixel coordinates to bar coordinates in the levels array
  const convertPixelCoordinatesToBarCoordinates = useCallback(
    (x: number, y: number, pixelsPerTick: number) => {
      const levelIndex = Math.floor(y / PIXELS_PER_LEVEL);
      const barIndex = getBarIndex(x, levels[levelIndex], pixelsPerTick, totalTicks, rangeMin);
      return { levelIndex, barIndex };
    },
    [getBarIndex, levels, totalTicks, rangeMin]
  );

  const render = useCallback(
    (pixelsPerTick: number) => {
      if (!levels.length) {
        return;
      }
      const ctx = graphRef.current?.getContext('2d')!;
      const graph = graphRef.current!;

      graph.height = PIXELS_PER_LEVEL * levels.length;
      graph.width = graph.clientWidth;
      ctx.textBaseline = 'middle';
      ctx.font = '13px Roboto, sans-serif';
      ctx.strokeStyle = 'white';

      for (let levelIndex = 0; levelIndex < levels.length; levelIndex++) {
        const level = levels[levelIndex];
        const dimensions = getRectDimensionsForLevel(level, levelIndex, totalTicks, rangeMin, pixelsPerTick);
        for (const rect of dimensions) {
          renderRect(ctx, rect, totalTicks, rangeMin, rangeMax, query, levelIndex, topLevelIndex);
        }
      }
    },
    [levels, query, rangeMax, rangeMin, topLevelIndex, totalTicks]
  );

  useEffect(() => {
    if (graphRef.current) {
      const pixelsPerTick = graphRef.current.clientWidth / totalTicks / (rangeMax - rangeMin);
      render(pixelsPerTick);

      graphRef.current.onclick = (e) => {
        const pixelsPerTick = graphRef.current!.clientWidth / totalTicks / (rangeMax - rangeMin);
        const { levelIndex, barIndex } = convertPixelCoordinatesToBarCoordinates(e.offsetX, e.offsetY, pixelsPerTick);
        if (barIndex === -1) {
          return;
        }
        if (!isNaN(levelIndex) && !isNaN(barIndex)) {
          setTopLevelIndex(levelIndex);
          setRangeMin(levels[levelIndex][barIndex].start / totalTicks);
          setRangeMax((levels[levelIndex][barIndex].start + levels[levelIndex][barIndex].value) / totalTicks);
        }
      };
    }
  }, [
    render,
    convertPixelCoordinatesToBarCoordinates,
    levels,
    rangeMin,
    rangeMax,
    topLevelIndex,
    totalTicks,
    windowWidth,
    setTopLevelIndex,
    setRangeMin,
    setRangeMax,
  ]);

  return <canvas className={styles.graph} ref={graphRef} data-testid="flamegraph" />;
};

const getStyles = () => ({
  graph: css`
    cursor: pointer;
    width: 100%;
  `,
});

// binary search for a bar in a level
const getBarIndex = (
  x: number,
  level: ItemWithStart[],
  pixelsPerTick: number,
  totalTicks: number,
  rangeMin: number
) => {
  if (level) {
    let start = 0;
    let end = level.length - 1;

    while (start <= end) {
      const midIndex = (start + end) >> 1;
      const startOfBar = getBarX(level[midIndex].start, totalTicks, rangeMin, pixelsPerTick);
      const startOfNextBar = getBarX(
        level[midIndex].start + level[midIndex].value,
        totalTicks,
        rangeMin,
        pixelsPerTick
      );

      if (startOfBar <= x && startOfNextBar >= x) {
        return startOfNextBar - startOfBar > COLLAPSE_THRESHOLD ? midIndex : -1;
      }

      if (startOfBar > x) {
        end = midIndex - 1;
      } else {
        start = midIndex + 1;
      }
    }
  }
  return -1;
};

export default FlameGraph;

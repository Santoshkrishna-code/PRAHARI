import React, { useEffect, useRef } from 'react';
import * as echarts from 'echarts';

interface EChartWrapperProps {
  options: echarts.EChartsOption;
  style?: React.CSSProperties;
}

export const EChartWrapper: React.FC<EChartWrapperProps> = ({ options, style = { height: '300px', width: '100%' } }) => {
  const chartRef = useRef<HTMLDivElement>(null);
  const chartInstance = useRef<echarts.ECharts | null>(null);

  useEffect(() => {
    if (chartRef.current) {
      chartInstance.current = echarts.init(chartRef.current);
      chartInstance.current.setOption(options);
    }

    const handleResize = () => {
      chartInstance.current?.resize();
    };

    window.addEventListener('resize', handleResize);

    return () => {
      window.removeEventListener('resize', handleResize);
      chartInstance.current?.dispose();
    };
  }, [options]);

  return <div ref={chartRef} style={style} className="border border-border rounded-lg p-2.5 bg-surface" />;
};

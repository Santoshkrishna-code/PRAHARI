import React from 'react';

export const Toolbar: React.FC<{ children: React.ReactNode; className?: string }> = ({ children, className }) => (
  <div className={`h-10 flex items-center gap-1.5 px-3 border-b border-white/[0.04] bg-white/[0.01] shrink-0 ${className || ''}`}>
    {children}
  </div>
);

export const ToolBtn: React.FC<{ children: React.ReactNode; active?: boolean; onClick?: () => void; className?: string }> = ({ children, active, onClick, className }) => (
  <button
    onClick={onClick}
    className={`h-7 px-2 rounded-md text-[11px] font-medium flex items-center gap-1.5 transition-colors ${
      active ? 'bg-indigo-500/15 text-indigo-400' : 'text-zinc-500 hover:text-zinc-300 hover:bg-white/[0.04]'
    } ${className || ''}`}
  >
    {children}
  </button>
);

export const ToolSep = () => <div className="w-px h-4 bg-white/[0.06] mx-1" />;

export const Chart: React.FC<{ data: number[]; color: string; threshold?: number; h?: number; fill?: number }> = ({ data, color, threshold, h = 140, fill = 0.12 }) => {
  const w = 800;
  const mx = Math.max(...data, threshold || 0) * 1.06;
  const mn = Math.min(...data) * 0.94;
  const y = (v: number) => h - 16 - ((v - mn) / (mx - mn)) * (h - 32);
  const x = (i: number) => (i / (data.length - 1)) * w;
  const pts = data.map((v, i) => `${x(i)},${y(v)}`).join(' ');
  return (
    <svg viewBox={`0 0 ${w} ${h}`} className="w-full" style={{ height: h }} preserveAspectRatio="none">
      <defs>
        <linearGradient id={`g${color.slice(1)}`} x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stopColor={color} stopOpacity={fill} />
          <stop offset="100%" stopColor={color} stopOpacity={0} />
        </linearGradient>
      </defs>
      {[0.25, 0.5, 0.75].map(f => (
        <line key={f} x1={0} y1={16 + f * (h - 32)} x2={w} y2={16 + f * (h - 32)} stroke="rgba(255,255,255,0.03)" />
      ))}
      {threshold && (
        <>
          <line x1={0} y1={y(threshold)} x2={w} y2={y(threshold)} stroke="#ef4444" strokeWidth="1" strokeDasharray="6,4" opacity="0.4" />
          <text x={w - 4} y={y(threshold) - 4} fill="#ef4444" fontSize="9" textAnchor="end" opacity="0.6">{threshold}</text>
        </>
      )}
      <polygon points={`${x(0)},${h - 16} ${pts} ${x(data.length - 1)},${h - 16}`} fill={`url(#g${color.slice(1)})`} />
      <polyline points={pts} fill="none" stroke={color} strokeWidth="1.5" strokeLinejoin="round" />
      <circle cx={x(data.length - 1)} cy={y(data[data.length - 1])} r="3" fill={color} />
      <circle cx={x(data.length - 1)} cy={y(data[data.length - 1])} r="7" fill={color} opacity="0.15" />
    </svg>
  );
};

export const Metric: React.FC<{ label: string; value: string | number; unit?: string; accent?: string; small?: boolean }> = ({ label, value, unit, accent, small }) => (
  <div>
    <p className="text-[10px] text-zinc-500 uppercase tracking-wider mb-0.5">{label}</p>
    <p className={`${small ? 'text-lg' : 'text-2xl'} font-semibold tracking-tight ${accent || 'text-white'}`}>
      {value}
      {unit && <span className="text-xs text-zinc-500 ml-0.5 font-normal">{unit}</span>}
    </p>
  </div>
);

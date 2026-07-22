import React from 'react';
import { X, ExternalLink, Activity, Shield, AlertTriangle, Wrench, Eye, GitBranch } from 'lucide-react';
import { knowledgeGraph, GraphEntity } from '../../services/knowledgeGraph';

interface EntityInspectorProps {
  entityId: string | null;
  onClose: () => void;
  onSelectEntity: (id: string) => void;
}

export const EntityInspector: React.FC<EntityInspectorProps> = ({ entityId, onClose, onSelectEntity }) => {
  if (!entityId) return null;

  const entity = knowledgeGraph.getEntity(entityId);
  const related = knowledgeGraph.getRelated(entityId);

  return (
    <div className="fixed inset-y-0 right-0 w-[440px] bg-zinc-900 border-l border-white/[0.08] shadow-2xl z-50 flex flex-col backdrop-blur-xl">
      {/* Header */}
      <div className="h-14 px-5 border-b border-white/[0.06] flex items-center justify-between bg-[#09090d]">
        <div>
          <span className="text-[10px] text-zinc-500 uppercase tracking-wider font-semibold">Entity Knowledge Graph</span>
          <h3 className="font-bold text-white text-base truncate max-w-[340px]">{entity?.name || entityId}</h3>
        </div>
        <button onClick={onClose} className="text-zinc-500 hover:text-white transition-colors">
          <X size={18} />
        </button>
      </div>

      <div className="flex-1 overflow-y-auto p-5 space-y-6">
        {/* Status Badge & Summary */}
        <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04] space-y-2">
          <div className="flex justify-between items-center">
            <span className="text-xs text-zinc-400 font-mono">{entity?.id}</span>
            <span className={`text-[10px] font-bold px-2 py-0.5 rounded-full ${
              entity?.status === 'Warning' ? 'bg-amber-500/15 text-amber-400' : 'bg-emerald-500/10 text-emerald-400'
            }`}>
              {entity?.status}
            </span>
          </div>
          <p className="text-sm font-medium text-white">{entity?.details}</p>
          <div className="text-xs text-zinc-500">Category Type: <strong className="text-zinc-300">{entity?.type}</strong></div>
        </div>

        {/* Connected Graph Relationships */}
        <div className="space-y-3">
          <div className="flex items-center gap-2">
            <GitBranch size={14} className="text-indigo-400" />
            <h4 className="text-xs font-bold text-zinc-400 uppercase tracking-wider">Connected Knowledge Graph Nodes ({related.length})</h4>
          </div>

          <div className="space-y-2">
            {related.map(rel => (
              <div
                key={rel.id}
                onClick={() => onSelectEntity(rel.id)}
                className="p-3 rounded-xl bg-white/[0.02] hover:bg-white/[0.06] border border-white/[0.04] cursor-pointer transition-colors flex items-center justify-between group"
              >
                <div className="space-y-0.5">
                  <div className="flex items-center gap-2">
                    <span className="text-xs font-semibold text-white group-hover:text-indigo-300 transition-colors">{rel.name}</span>
                    <span className="text-[9px] px-1.5 py-0.5 rounded bg-zinc-800 text-zinc-400">{rel.type}</span>
                  </div>
                  <p className="text-[11px] text-zinc-500">{rel.details}</p>
                </div>
                <ExternalLink size={14} className="text-zinc-600 group-hover:text-indigo-400 transition-colors shrink-0" />
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

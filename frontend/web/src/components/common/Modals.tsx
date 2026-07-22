import React, { useState } from 'react';
import { X, Plus, AlertTriangle, Shield, Package, UserPlus } from 'lucide-react';
import { Asset, Incident } from '../../types';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  children: React.ReactNode;
}

export const Modal: React.FC<ModalProps> = ({ isOpen, onClose, title, children }) => {
  if (!isOpen) return null;
  return (
    <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div className="w-full max-w-lg bg-zinc-900 border border-white/[0.08] rounded-2xl shadow-2xl overflow-hidden flex flex-col max-h-[90vh]">
        <div className="h-12 px-5 border-b border-white/[0.06] flex items-center justify-between bg-white/[0.01]">
          <h3 className="font-bold text-white text-sm">{title}</h3>
          <button onClick={onClose} className="text-zinc-500 hover:text-white transition-colors">
            <X size={16} />
          </button>
        </div>
        <div className="p-5 overflow-y-auto flex-1">{children}</div>
      </div>
    </div>
  );
};

export const CreateAssetModal: React.FC<{ isOpen: boolean; onClose: () => void; onCreate: (asset: Asset) => void }> = ({
  isOpen,
  onClose,
  onCreate,
}) => {
  const [name, setName] = useState('');
  const [loc, setLoc] = useState('DC-101 Recirc');
  const [type, setType] = useState('Centrifugal Pump');
  const [owner, setOwner] = useState('Mechanical Team');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onCreate({
      name: name || 'New Asset P-105',
      loc: loc,
      type: type,
      health: 100,
      rul: '120d',
      st: 'Running',
      owner: owner,
      vib: 1.2,
      temp: 42,
    });
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Register New Equipment Asset">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Asset Tag Name</label>
          <input
            type="text"
            required
            value={name}
            onChange={e => setName(e.target.value)}
            placeholder="e.g. Pump P-105"
            className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
          />
        </div>
        <div className="grid grid-cols-2 gap-3">
          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Location / Zone</label>
            <input
              type="text"
              value={loc}
              onChange={e => setLoc(e.target.value)}
              className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
            />
          </div>
          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Equipment Category</label>
            <input
              type="text"
              value={type}
              onChange={e => setType(e.target.value)}
              className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
            />
          </div>
        </div>
        <div>
          <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Responsible Owner Team</label>
          <input
            type="text"
            value={owner}
            onChange={e => setOwner(e.target.value)}
            className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
          />
        </div>
        <button type="submit" className="w-full py-2.5 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-xs transition-all">
          Save & Register Asset
        </button>
      </form>
    </Modal>
  );
};

export const ReportIncidentModal: React.FC<{ isOpen: boolean; onClose: () => void; onReport: (inc: Incident) => void }> = ({
  isOpen,
  onClose,
  onReport,
}) => {
  const [title, setTitle] = useState('');
  const [asset, setAsset] = useState('Pump P-102');
  const [sev, setSev] = useState('Warning');
  const [desc, setDesc] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onReport({
      id: `INC-2026-${Math.floor(1000 + Math.random() * 9000)}`,
      title: title || 'Unplanned Vibration Excursion',
      desc: desc || 'Vibration probe threshold exceeded during high-pressure run.',
      sev: sev,
      asset: asset,
      st: 'Under Investigation',
      time: 'Just Now',
    });
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Report New Incident / Hazard">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Incident Summary Title</label>
          <input
            type="text"
            required
            value={title}
            onChange={e => setTitle(e.target.value)}
            placeholder="e.g. Bearing overheating on Pump P-102"
            className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
          />
        </div>
        <div className="grid grid-cols-2 gap-3">
          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Associated Asset</label>
            <input
              type="text"
              value={asset}
              onChange={e => setAsset(e.target.value)}
              className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
            />
          </div>
          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Severity Level</label>
            <select
              value={sev}
              onChange={e => setSev(e.target.value)}
              className="w-full px-3 py-2 rounded-xl bg-zinc-800 border border-white/[0.08] text-xs text-white"
            >
              <option value="Critical">Critical (P1)</option>
              <option value="Warning">Warning (P2)</option>
              <option value="Info">Low (P3)</option>
            </select>
          </div>
        </div>
        <div>
          <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">Detailed Event Description</label>
          <textarea
            rows={3}
            value={desc}
            onChange={e => setDesc(e.target.value)}
            placeholder="Provide context on initial observations, sensor values, or PPE violations..."
            className="w-full px-3 py-2 rounded-xl bg-white/[0.03] border border-white/[0.08] text-xs text-white"
          />
        </div>
        <button type="submit" className="w-full py-2.5 rounded-xl bg-red-600 hover:bg-red-500 text-white font-semibold text-xs transition-all">
          Dispatch Incident Investigation
        </button>
      </form>
    </Modal>
  );
};

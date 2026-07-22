import React, { useState } from 'react';
import { Sparkles, X, ChevronRight, CheckCircle2, Zap } from 'lucide-react';
import { PageId } from '../../types';

interface GuidedTourProps {
  onNavigate: (pageId: PageId) => void;
  onDismiss: () => void;
}

export const GuidedTour: React.FC<GuidedTourProps> = ({ onNavigate, onDismiss }) => {
  const [tourStep, setTourStep] = useState(0);

  const steps: { title: string; desc: string; page: PageId; actionText: string }[] = [
    {
      title: 'Welcome to PRAHARI Command Center',
      desc: 'This is your executive operational dashboard monitoring live signals, active risks, and AI agent actions.',
      page: 'command-center',
      actionText: 'Explore Industrial Twin →',
    },
    {
      title: 'CAD / SCADA Industrial Digital Twin',
      desc: 'Click on any plant equipment (Reactor B, Pump P-102, HX-04) to overlay sensor trends and failure mode predictions.',
      page: 'industrial-twin',
      actionText: 'Open AI Command Center →',
    },
    {
      title: 'Autonomous AI Supervisor Engine',
      desc: 'Watch 10 domain agents autonomously compute 5-Whys root cause analysis, evidence correlation, and ISO compliance.',
      page: 'ai-command',
      actionText: 'View Edge Vision Detections →',
    },
    {
      title: 'YOLOv8 Edge Computer Vision',
      desc: 'Inspect live camera streams, PPE helmet violations, Jetson AGX node frame rates, and latency counters.',
      page: 'vision-intel',
      actionText: 'Finish Product Tour',
    },
  ];

  const current = steps[tourStep];

  const handleNext = () => {
    if (tourStep < steps.length - 1) {
      const nextIdx = tourStep + 1;
      setTourStep(nextIdx);
      onNavigate(steps[nextIdx].page);
    } else {
      onDismiss();
    }
  };

  return (
    <div className="bg-indigo-950/90 border-b border-indigo-500/30 px-6 py-3 flex items-center justify-between z-40 text-xs shadow-lg backdrop-blur-md">
      <div className="flex items-center gap-3">
        <div className="w-7 h-7 rounded-lg bg-indigo-600 flex items-center justify-center shrink-0">
          <Sparkles size={14} className="text-white" />
        </div>
        <div>
          <span className="font-bold text-white tracking-wide">
            Interactive Product Tour ({tourStep + 1}/{steps.length}): {current.title}
          </span>
          <p className="text-zinc-300 text-[11px] mt-0.5">{current.desc}</p>
        </div>
      </div>

      <div className="flex items-center gap-3">
        <button
          onClick={handleNext}
          className="px-3 py-1.5 rounded-lg bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-xs shadow transition-all flex items-center gap-1"
        >
          {current.actionText}
        </button>
        <button
          onClick={onDismiss}
          className="text-zinc-400 hover:text-white transition-colors p-1"
        >
          <X size={16} />
        </button>
      </div>
    </div>
  );
};

import React, { useState } from 'react';
import { Brain, X, Sparkles, Send, CheckCircle2, ArrowRight, MessageSquare } from 'lucide-react';
import { PageId } from '../../types';

interface AIDrawerProps {
  isOpen: boolean;
  onClose: () => void;
  currentPage: PageId;
  selectedEntityId?: string;
}

export const AIDrawer: React.FC<AIDrawerProps> = ({ isOpen, onClose, currentPage, selectedEntityId }) => {
  const [messages, setMessages] = useState<{ sender: 'user' | 'ai'; text: string; time: string }[]>([
    {
      sender: 'ai',
      text: `Hello! I am the PRAHARI Autonomous AI Safety Supervisor. I am currently monitoring workspace context: **${currentPage}** ${
        selectedEntityId ? `(Entity: ${selectedEntityId})` : ''
      }. How can I assist your safety investigation?`,
      time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
    },
  ]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);

  if (!isOpen) return null;

  const getSuggestedPrompts = () => {
    switch (currentPage) {
      case 'assets':
      case 'industrial-twin':
        return ['Explain Pump P-102 bearing wear trajectory', 'Predict MTBF for Heat Exchanger HX-04', 'Audit PM overdue history for P-102'];
      case 'incidents':
        return ['Synthesize 5-Whys root cause summary', 'Draft OSHA incident report for INC-0447', 'Verify CAPA completion SLA'];
      case 'permits':
        return ['Audit PTW-8902 LOTO isolation compliance', 'Check gas test threshold safety margin', 'Verify contractor isolation tags'];
      case 'vision-intel':
        return ['Summarize CAM-002 PPE violation frequency', 'Verify Jetson AGX node frame rates', 'Generate perimeter security report'];
      default:
        return ['Summarize plant safety status', 'Correlate vibration anomaly with PTW permits', 'Draft executive ISO 45001 compliance summary'];
    }
  };

  const handleSend = (queryText?: string) => {
    const textToSend = queryText || input;
    if (!textToSend.trim()) return;

    const userMsg = {
      sender: 'user' as const,
      text: textToSend,
      time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
    };

    setMessages(prev => [...prev, userMsg]);
    if (!queryText) setInput('');
    setLoading(true);

    setTimeout(() => {
      let responseText = 'AI Supervisor Reasoning: Evaluated real-time event stream and knowledge graph relations. All safety constraints validated.';

      if (textToSend.toLowerCase().includes('bearing') || textToSend.toLowerCase().includes('p-102') || textToSend.toLowerCase().includes('vibration')) {
        responseText = 'Pump P-102 analysis: Vibration velocity reading 11.8 mm/s. Physics-Informed Neural Network predicts bearing race degradation (72% probability). Remaining useful life: 18 days. Root cause: Scheduled PM interval was exceeded by 14 days.';
      } else if (textToSend.toLowerCase().includes('5-whys') || textToSend.toLowerCase().includes('incident') || textToSend.toLowerCase().includes('root cause')) {
        responseText = 'INC-2026-0447 Root Cause Synthesis: 1. Vibration reached 11.8 mm/s → 2. Bearing race misalignment → 3. Lubrication breakdown from overheating → 4. PM interval exceeded by 14 days due to un-escalated WO-7821.';
      } else if (textToSend.toLowerCase().includes('permit') || textToSend.toLowerCase().includes('loto') || textToSend.toLowerCase().includes('ptw')) {
        responseText = 'Permit PTW-8902 Audit: Hot Work approved for Tank T-204. Physical isolation lock verified on Valve V-88. Gas test readings: O₂ 20.9%, H₂S 0ppm. Zero permit conflicts detected.';
      }

      setMessages(prev => [
        ...prev,
        {
          sender: 'ai',
          text: responseText,
          time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
        },
      ]);
      setLoading(false);
    }, 700);
  };

  return (
    <div className="fixed inset-y-0 right-0 w-[420px] bg-zinc-900 border-l border-white/[0.08] shadow-2xl z-50 flex flex-col backdrop-blur-xl">
      {/* Header */}
      <div className="h-14 px-5 border-b border-white/[0.06] flex items-center justify-between bg-[#09090d]">
        <div className="flex items-center gap-2.5">
          <div className="w-7 h-7 rounded-lg bg-indigo-600 flex items-center justify-center">
            <Brain size={15} className="text-white" />
          </div>
          <div>
            <h3 className="font-bold text-white text-sm">Ask AI Assistant</h3>
            <p className="text-[10px] text-zinc-500">Context: {currentPage} Workspace</p>
          </div>
        </div>
        <button onClick={onClose} className="text-zinc-500 hover:text-white transition-colors">
          <X size={18} />
        </button>
      </div>

      {/* Messages Feed */}
      <div className="flex-1 overflow-y-auto p-4 space-y-3.5">
        {messages.map((m, i) => (
          <div
            key={i}
            className={`flex flex-col ${m.sender === 'user' ? 'items-end' : 'items-start'}`}
          >
            <div
              className={`max-w-[90%] p-3 rounded-2xl text-xs leading-relaxed ${
                m.sender === 'user'
                  ? 'bg-indigo-600 text-white rounded-br-none'
                  : 'bg-zinc-800/80 border border-white/[0.06] text-zinc-200 rounded-bl-none'
              }`}
            >
              {m.text}
            </div>
            <span className="text-[9px] text-zinc-600 mt-1 px-1">{m.time}</span>
          </div>
        ))}
        {loading && (
          <div className="flex items-center gap-2 text-xs text-indigo-400 p-2">
            <Sparkles size={14} className="animate-spin" />
            <span>Reasoning across Knowledge Graph...</span>
          </div>
        )}
      </div>

      {/* Suggested Prompts */}
      <div className="px-4 py-2 border-t border-white/[0.04] bg-white/[0.01]">
        <p className="text-[9px] font-semibold text-zinc-500 uppercase tracking-wider mb-1.5">Contextual AI Prompts</p>
        <div className="flex flex-wrap gap-1.5">
          {getSuggestedPrompts().map(prompt => (
            <button
              key={prompt}
              onClick={() => handleSend(prompt)}
              className="text-[10px] px-2 py-1 rounded-md bg-white/[0.03] hover:bg-indigo-600/20 hover:text-indigo-300 border border-white/[0.06] text-zinc-400 text-left transition-all"
            >
              {prompt}
            </button>
          ))}
        </div>
      </div>

      {/* Input */}
      <div className="p-3 border-t border-white/[0.06] bg-[#09090d]">
        <form
          onSubmit={e => {
            e.preventDefault();
            handleSend();
          }}
          className="flex items-center gap-2"
        >
          <input
            type="text"
            value={input}
            onChange={e => setInput(e.target.value)}
            placeholder="Ask AI about safety, RUL, or RCA..."
            className="flex-1 px-3 py-2 rounded-xl bg-white/[0.04] border border-white/[0.08] text-xs text-white placeholder-zinc-500 focus:outline-none focus:border-indigo-500"
          />
          <button
            type="submit"
            className="p-2 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white transition-colors"
          >
            <Send size={14} />
          </button>
        </form>
      </div>
    </div>
  );
};

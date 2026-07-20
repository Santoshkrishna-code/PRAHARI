import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Input } from '@prahari/components/Input.tsx';

interface Message {
  role: 'USER' | 'COPILOT';
  content: string;
}

export const AICopilot: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([
    { role: 'COPILOT', content: 'Hello! I am your safety regulations RCA assistant. Ask me anything about standard LOTO procedures or Chemical SDS bounds.' }
  ]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSend = () => {
    if (!input.trim()) return;
    const userMsg: Message = { role: 'USER', content: input };
    setMessages((prev) => [...prev, userMsg]);
    setInput('');
    setLoading(true);

    setTimeout(() => {
      const copilotMsg: Message = {
        role: 'COPILOT',
        content: `Based on chemical safety rules, visual spill alarms require immediate pipeline isolation. Verified checklist permits: PRM-001.`
      };
      setMessages((prev) => [...prev, copilotMsg]);
      setLoading(false);
    }, 400);
  };

  return (
    <div className="flex flex-col h-[70vh] border border-border rounded-lg bg-surface overflow-hidden w-full">
      <div className="bg-primary text-background px-4 py-3 text-sm font-bold flex items-center justify-between">
        <span>PRAHARI AI Copilot Chat Workspace</span>
      </div>

      <div className="flex-1 overflow-y-auto p-4 flex flex-col gap-3">
        {messages.map((m, idx) => (
          <div
            key={idx}
            className={`max-w-[80%] p-3 rounded-lg text-sm leading-relaxed ${
              m.role === 'USER'
                ? 'self-end bg-primary text-background'
                : 'self-start bg-background text-text border border-border'
            }`}
          >
            {m.content}
          </div>
        ))}
        {loading && <div className="self-start text-xs text-muted">Copilot is thinking...</div>}
      </div>

      <div className="border-t border-border p-3 flex gap-2 bg-background">
        <Input
          placeholder="Ask safety regulations copilot..."
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => e.key === 'Enter' && handleSend()}
        />
        <Button onClick={handleSend} isLoading={loading}>Send</Button>
      </div>
    </div>
  );
};

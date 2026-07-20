import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Input } from '@prahari/components/Input.tsx';

interface Message {
  role: 'USER' | 'COPILOT';
  content: string;
}

export const AICopilot: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([
    { role: 'COPILOT', content: 'Hello! Ask me standard EHS regulations or SDS safety codes.' }
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
      setMessages((prev) => [
        ...prev,
        { role: 'COPILOT', content: 'Safety procedure 02 requires double isolation checks for hot welding permits.' }
      ]);
      setLoading(false);
    }, 400);
  };

  return (
    <div className="flex flex-col h-[60vh] border border-border rounded-xl bg-surface overflow-hidden w-full max-w-lg mx-auto shadow-sm">
      <div className="bg-primary text-background px-4 py-3 text-sm font-bold">
        <span>Field AI assistant Copilot</span>
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
      </div>

      <div className="border-t border-border p-3 flex gap-2 bg-background">
        <Input
          placeholder="Ask safety regulations..."
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => e.key === 'Enter' && handleSend()}
        />
        <Button onClick={handleSend} isLoading={loading}>Send</Button>
      </div>
    </div>
  );
};

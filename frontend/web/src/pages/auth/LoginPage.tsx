import React, { useState } from 'react';
import { Shield, Zap, Lock, Mail, User, CheckCircle2, ArrowRight, KeyRound } from 'lucide-react';
import { UserRole } from '../../types/auth';

interface LoginPageProps {
  onLoginSuccess: (user: { name: string; email: string; role: UserRole; orgName: string; plantName: string }) => void;
  onNavigateRegister: () => void;
  onNavigateLanding: () => void;
}

export const LoginPage: React.FC<LoginPageProps> = ({ onLoginSuccess, onNavigateRegister, onNavigateLanding }) => {
  const [email, setEmail] = useState('santosh@prahari-industrial.io');
  const [password, setPassword] = useState('••••••••••••');
  const [selectedRole, setSelectedRole] = useState<UserRole>('Plant Manager');
  const [loading, setLoading] = useState(false);
  const [forgotOpen, setForgotOpen] = useState(false);
  const [forgotSent, setForgotSent] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      onLoginSuccess({
        name: email.split('@')[0].replace('.', ' ').toUpperCase(),
        email: email,
        role: selectedRole,
        orgName: 'Alpha Chemical Refinery Inc.',
        plantName: 'Plant Alpha (Gulf Coast)',
      });
    }, 600);
  };

  const handleQuickLogin = (role: UserRole) => {
    setSelectedRole(role);
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      onLoginSuccess({
        name: `${role} Lead`,
        email: `${role.toLowerCase().replace(/\s+/g, '.')}@prahari-industrial.io`,
        role: role,
        orgName: 'Alpha Chemical Refinery Inc.',
        plantName: 'Plant Alpha (Gulf Coast)',
      });
    }, 400);
  };

  return (
    <div className="min-h-screen bg-[#06080d] text-zinc-100 flex flex-col justify-center items-center px-4 font-sans relative overflow-hidden">
      {/* Background Radial Glow */}
      <div className="absolute top-1/4 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-indigo-600/10 rounded-full blur-[140px] pointer-events-none" />

      {/* Brand Header */}
      <div className="mb-8 text-center cursor-pointer" onClick={onNavigateLanding}>
        <div className="inline-flex items-center gap-2 mb-2">
          <div className="w-9 h-9 rounded-xl bg-indigo-600 flex items-center justify-center shadow-lg shadow-indigo-500/30">
            <Zap size={18} className="text-white" />
          </div>
          <span className="font-display text-2xl font-bold text-white tracking-tight">PRAHARI</span>
        </div>
        <p className="text-xs text-zinc-500">Enterprise AI Safety Operating System</p>
      </div>

      {/* Login Card */}
      <div className="w-full max-w-md bg-zinc-900/80 border border-white/[0.08] rounded-2xl p-8 shadow-2xl backdrop-blur-xl space-y-6">
        <div className="space-y-1 text-center">
          <h2 className="text-xl font-bold text-white">Sign In to Your Workspace</h2>
          <p className="text-xs text-zinc-400">Enter your enterprise credentials or pick a demo persona</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1.5">
              Work Email
            </label>
            <div className="relative">
              <Mail size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-zinc-500" />
              <input
                type="email"
                required
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="w-full pl-9 pr-4 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white placeholder-zinc-600 focus:outline-none focus:border-indigo-500 transition-colors"
                placeholder="name@company.com"
              />
            </div>
          </div>

          <div>
            <div className="flex justify-between items-center mb-1.5">
              <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider">
                Password
              </label>
              <button
                type="button"
                onClick={() => setForgotOpen(true)}
                className="text-xs text-indigo-400 hover:underline"
              >
                Forgot Password?
              </button>
            </div>
            <div className="relative">
              <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-zinc-500" />
              <input
                type="password"
                required
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full pl-9 pr-4 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white focus:outline-none focus:border-indigo-500 transition-colors"
              />
            </div>
          </div>

          <div>
            <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1.5">
              Assigned Role Scope
            </label>
            <select
              value={selectedRole}
              onChange={e => setSelectedRole(e.target.value as UserRole)}
              className="w-full px-3 py-2.5 rounded-xl bg-zinc-800 border border-white/[0.08] text-sm text-zinc-200 focus:outline-none focus:border-indigo-500"
            >
              <option value="Plant Manager">Plant Manager</option>
              <option value="Safety Officer">Safety Officer</option>
              <option value="Reliability Engineer">Reliability Engineer</option>
              <option value="Maintenance Engineer">Maintenance Engineer</option>
              <option value="Operator">Operator</option>
              <option value="Administrator">Administrator</option>
              <option value="Auditor">Auditor</option>
              <option value="Executive">Executive</option>
            </select>
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-3 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-sm shadow-lg shadow-indigo-600/30 transition-all flex items-center justify-center gap-2"
          >
            {loading ? (
              <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin" />
            ) : (
              <>
                Sign In to Platform <ArrowRight size={16} />
              </>
            )}
          </button>
        </form>

        {/* Quick Demo Login Preset Buttons */}
        <div className="pt-2 border-t border-white/[0.06] space-y-2">
          <p className="text-[10px] text-zinc-500 text-center uppercase tracking-wider font-semibold">
            One-Click Persona Login
          </p>
          <div className="grid grid-cols-2 gap-2">
            {(['Plant Manager', 'Safety Officer', 'Reliability Engineer', 'Administrator'] as UserRole[]).map(r => (
              <button
                key={r}
                onClick={() => handleQuickLogin(r)}
                className="px-2.5 py-1.5 rounded-lg bg-white/[0.03] hover:bg-white/[0.08] border border-white/[0.06] text-xs text-zinc-300 transition-colors flex items-center gap-1.5 justify-center"
              >
                <User size={12} className="text-indigo-400" /> {r}
              </button>
            ))}
          </div>
        </div>

        <div className="text-center text-xs text-zinc-500">
          Need a new organization tenant?{' '}
          <button onClick={onNavigateRegister} className="text-indigo-400 font-semibold hover:underline">
            Register Here
          </button>
        </div>
      </div>

      {/* Forgot Password Modal */}
      {forgotOpen && (
        <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center p-4">
          <div className="w-full max-w-sm bg-zinc-900 border border-white/[0.08] rounded-2xl p-6 space-y-4 shadow-2xl">
            <div className="flex items-center gap-2 text-indigo-400">
              <KeyRound size={20} />
              <h3 className="font-bold text-white text-base">Reset Password</h3>
            </div>
            {forgotSent ? (
              <div className="space-y-3 text-center">
                <CheckCircle2 size={32} className="text-emerald-400 mx-auto" />
                <p className="text-sm text-zinc-300">Reset link sent to <strong>{email}</strong>!</p>
                <button
                  onClick={() => { setForgotOpen(false); setForgotSent(false); }}
                  className="w-full py-2 rounded-xl bg-zinc-800 text-white text-xs font-semibold"
                >
                  Return to Sign In
                </button>
              </div>
            ) : (
              <div className="space-y-3">
                <p className="text-xs text-zinc-400">Enter your registered email address to receive password reset instructions.</p>
                <input
                  type="email"
                  value={email}
                  onChange={e => setEmail(e.target.value)}
                  className="w-full px-3 py-2 rounded-xl bg-white/[0.04] border border-white/[0.08] text-xs text-white"
                />
                <div className="flex gap-2">
                  <button
                    onClick={() => setForgotOpen(false)}
                    className="flex-1 py-2 rounded-xl bg-zinc-800 text-zinc-400 text-xs font-medium"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={() => setForgotSent(true)}
                    className="flex-1 py-2 rounded-xl bg-indigo-600 text-white text-xs font-semibold"
                  >
                    Send Email
                  </button>
                </div>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

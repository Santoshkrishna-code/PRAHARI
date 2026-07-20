import { create } from 'zustand';

interface PlatformState {
  sidebarOpen: boolean;
  toggleSidebar: () => void;
  onlineStatus: boolean;
  setOnlineStatus: (status: boolean) => void;
}

export const usePlatformStore = create<PlatformState>((set) => ({
  sidebarOpen: true,
  toggleSidebar: () => set((state) => ({ sidebarOpen: !state.sidebarOpen })),
  onlineStatus: typeof navigator !== 'undefined' ? navigator.onLine : true,
  setOnlineStatus: (status) => set({ onlineStatus: status })
}));

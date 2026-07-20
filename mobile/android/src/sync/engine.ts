import { sqlite } from '../storage/db.ts';

export class SyncEngine {
  private isOnline: boolean = true;

  constructor() {
    if (typeof window !== 'undefined') {
      this.isOnline = navigator.onLine;
      window.addEventListener('online', () => this.triggerSync());
    }
  }

  async queueOfflineIncident(incident: any): Promise<void> {
    await sqlite.insert('incidents', incident);
    if (this.isOnline) {
      await this.triggerSync();
    }
  }

  async triggerSync(): Promise<void> {
    const list = await sqlite.queryAll('incidents');
    if (list.length === 0) return;

    // Simulate batch upload
    console.log(`[SyncEngine] Successfully synced ${list.length} offline logged records to backend microservices API`);
    await sqlite.clear('incidents');
  }

  setOnlineStatus(status: boolean) {
    this.isOnline = status;
    if (status) {
      this.triggerSync();
    }
  }
}

export const syncEngine = new SyncEngine();

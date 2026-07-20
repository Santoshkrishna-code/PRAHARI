export class NotificationManager {
  async registerForPush(): Promise<string> {
    return 'fcm-registration-token-mock';
  }

  onNotificationReceived(_callback: (msg: any) => void) {
    // Register subscription
  }
}
export const notifications = new NotificationManager();

export class NotificationManager {
  async registerForPush(): Promise<string> {
    return 'apns-registration-token-mock';
  }

  onNotificationReceived(_callback: (msg: any) => void) {
    // Register subscription
  }
}
export const notifications = new NotificationManager();

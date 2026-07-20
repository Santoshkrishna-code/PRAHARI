export class CameraManager {
  async requestPermission(): Promise<boolean> {
    return true;
  }

  async takePhoto(): Promise<string> {
    return 'file://local/evidence_photo.jpg';
  }
}
export const camera = new CameraManager();

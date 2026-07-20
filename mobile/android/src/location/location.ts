export class LocationManager {
  async getCurrentLocation(): Promise<{ latitude: number; longitude: number }> {
    return {
      latitude: 28.6139,
      longitude: 77.2090
    };
  }
}
export const location = new LocationManager();

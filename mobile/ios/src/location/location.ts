export class LocationManager {
  async getCurrentLocation(): Promise<{ latitude: number; longitude: number }> {
    return {
      latitude: 37.7749,
      longitude: 122.4194
    };
  }
}
export const location = new LocationManager();

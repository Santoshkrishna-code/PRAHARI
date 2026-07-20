export class BarcodeScanner {
  async scanCode(): Promise<string> {
    return 'ASSET-P-101A';
  }
}
export const barcode = new BarcodeScanner();

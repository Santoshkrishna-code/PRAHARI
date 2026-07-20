export class SQLiteMock {
  private localRecords: Map<string, any[]> = new Map();

  constructor() {
    this.localRecords.set('incidents', []);
    this.localRecords.set('permits', []);
  }

  async insert(table: string, record: any): Promise<void> {
    const list = this.localRecords.get(table) || [];
    list.push(record);
    this.localRecords.set(table, list);
  }

  async queryAll(table: string): Promise<any[]> {
    return this.localRecords.get(table) || [];
  }

  async clear(table: string): Promise<void> {
    this.localRecords.set(table, []);
  }
}

export const sqlite = new SQLiteMock();
export const secureKeychain = {
  getToken: () => localStorage.getItem('prahari-jwt-token'),
  setToken: (tok: string) => localStorage.setItem('prahari-jwt-token', tok)
};

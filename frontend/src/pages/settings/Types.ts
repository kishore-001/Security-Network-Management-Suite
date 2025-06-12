export interface User {
  id: number;
  username: string;
  role: string;
  email: string;
}

export interface MACAddress {
  id: number;
  address: string;
  type: 'whitelist' | 'blacklist';
}

// auth/auth.ts
class AuthService {
  private static readonly ACCESS_TOKEN_KEY = 'access_token';
  private static readonly BACKEND_URL = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8000';



  // Get user role from access token
  static getUserRole(): 'admin' | 'viewer' | null {
    try {
      const token = this.getAccessToken();
      if (!token) return null;

      const payload = JSON.parse(atob(token.split('.')[1]));
      return payload.role || null;
    } catch (error) {
      console.error('Error parsing token for role:', error);
      return null;
    }
  }

  // Get user info from token
  static getUserInfo(): { username?: string; role?: 'admin' | 'viewer' } | null {
    try {
      const token = this.getAccessToken();
      if (!token) return null;

      const payload = JSON.parse(atob(token.split('.')[1]));
      return {
        username: payload.username || payload.sub,
        role: payload.role
      };
    } catch (error) {
      console.error('Error parsing token for user info:', error);
      return null;
    }
  }

  // Check if user has admin role
  static isAdmin(): boolean {
    return this.getUserRole() === 'admin';
  }

  // Check if user has viewer role
  static isViewer(): boolean {
    return this.getUserRole() === 'viewer';
  }

  // Check if user can access a specific route
  static canAccessRoute(path: string): boolean {
    const role = this.getUserRole();
    if (!role) return false;

    if (role === 'admin') return true;

    // Viewer can only access specific routes
    const viewerAllowedRoutes = ['/health', '/alert', '/log'];
    return viewerAllowedRoutes.includes(path);
  }
 

  // Check if user is authorized
  static async authorized(): Promise<boolean> {
    try {
      const accessToken = this.getAccessToken();
      
      // If no access token, try to refresh
      if (!accessToken) {
        console.log('No access token found, attempting to refresh...');
        return await this.attemptTokenRefresh();
      }

      // Check if access token is expired
      if (this.isTokenExpired(accessToken)) {
        console.log('Access token expired, attempting to refresh...');
        return await this.attemptTokenRefresh();
      }

      // Access token is valid
      console.log('Access token is valid');
      return true;

    } catch (error) {
      console.error('Authorization check failed:', error);
      return false;
    }
  }

  // Attempt to refresh the access token
  private static async attemptTokenRefresh(): Promise<boolean> {
    // First verify if refresh token is valid
    const isRefreshTokenValid = await this.verifyRefreshToken();
    
    if (!isRefreshTokenValid) {
      console.log('Refresh token is invalid or missing, redirecting to login');
      this.clearTokens();
      this.redirectToLogin();
      return false;
    }

    try {
      const response = await fetch(`${this.BACKEND_URL}/api/auth/refresh`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      if (response.ok) {
        const data = await response.json();
        
        if (data.status === 'ok' && data.access_token) {
          this.setAccessToken(data.access_token);
          console.log('Token refreshed successfully');
          return true;
        }
      }

      console.log('Token refresh failed, redirecting to login');
      this.clearTokens();
      this.redirectToLogin();
      return false;

    } catch (error) {
      console.error('Token refresh error:', error);
      this.clearTokens();
      this.redirectToLogin();
      return false;
    }
  }

  // Verify refresh token validity
  private static async verifyRefreshToken(): Promise<boolean> {
    try {
      const response = await fetch(`${this.BACKEND_URL}/api/auth/verify`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      if (response.ok) {
        const data = await response.json();
        return data.status === 'authorized';
      }
      return false;

    } catch (error) {
      console.error('Error verifying refresh token:', error);
      return false;
    }
  }

  // Get access token from localStorage
  static getAccessToken(): string | null {
    return localStorage.getItem(this.ACCESS_TOKEN_KEY);
  }

  // Set access token in localStorage
  static setAccessToken(token: string): void {
    localStorage.setItem(this.ACCESS_TOKEN_KEY, token);
  }

  // Clear access token from localStorage
  static clearTokens(): void {
    localStorage.removeItem(this.ACCESS_TOKEN_KEY);
  }

  // Check if token is expired
  private static isTokenExpired(token: string): boolean {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      const currentTime = Math.floor(Date.now() / 1000);
      return payload.exp < currentTime;
    } catch (error) {
      console.error('Error parsing token:', error);
      return true;
    }
  }

  // Redirect to login page
  public static redirectToLogin(): void {
    window.location.href = '/login';
  }

  // Get authorization header for API requests
  static getAuthHeader(): { Authorization: string } | Record<string, never> {
    const token = this.getAccessToken();
    return token ? { Authorization: `Bearer ${token}` } : {};
  }

  // Get complete headers with auth and content type
  static getApiHeaders(additionalHeaders: Record<string, string> = {}): Record<string, string> {
    const authHeaders = this.getAuthHeader();
    return {
      'Content-Type': 'application/json',
      ...authHeaders,
      ...additionalHeaders,
    };
  }

  // Make authenticated API request with automatic token refresh
  static async makeAuthenticatedRequest(
    url: string, 
    options: RequestInit = {}
  ): Promise<Response> {
    // Ensure we have a valid token
    const isAuthorized = await this.authorized();
    if (!isAuthorized) {
      throw new Error('Not authorized');
    }

    // Add auth headers to the request
    const authHeaders = this.getAuthHeader();
    const headers = {
      'Content-Type': 'application/json',
      ...authHeaders,
      ...(options.headers || {}),
    };

    const response = await fetch(url, {
      ...options,
      headers,
      credentials: 'include',
    });

    // If unauthorized, try to refresh token and retry once
    if (response.status === 401) {
      console.log('Request unauthorized, attempting token refresh...');
      const refreshed = await this.attemptTokenRefresh();
      
      if (refreshed) {
        // Retry with new token
        const newAuthHeaders = this.getAuthHeader();
        const retryHeaders = {
          'Content-Type': 'application/json',
          ...newAuthHeaders,
          ...(options.headers || {}),
        };

        return fetch(url, {
          ...options,
          headers: retryHeaders,
          credentials: 'include',
        });
      } else {
        throw new Error('Authentication failed');
      }
    }

    return response;
  }

  // Logout user
  static async logout(): Promise<void> {
    this.clearTokens();
    this.redirectToLogin();
  }
}

export default AuthService;

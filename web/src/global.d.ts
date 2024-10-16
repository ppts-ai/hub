export {};

declare global {
  interface Window {
    __TAURI__?: {
      core: {
        invoke: (command: string, args: object) => void;
      }
    }; 
  }
}
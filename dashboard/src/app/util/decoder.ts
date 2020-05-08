import {Injectable} from '@angular/core';

@Injectable({ providedIn: 'root' })
export class Decoder {
  // Decode a base64 string into a Uint8Array.
  _decodeBuffer(value: string) {
    return Uint8Array.from(atob(value), c => c.charCodeAt(0));
  }

  // Encode an ArrayBuffer into a base64 string.
  _encodeBuffer(value: Iterable<number>) {
    return btoa(new Uint8Array(value).reduce((s, byte) => s + String.fromCharCode(byte), ''));
  }
}

import { Injectable } from '@nestjs/common';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UtilsService {
  hashPassword(password: string): string {
    return bcrypt.hashSync(password, 12);
  }

  compareHash(password: string, hashed: string): boolean {
    return bcrypt.compareSync(password, hashed);
  }
}

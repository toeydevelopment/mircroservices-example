import {
  CACHE_MANAGER,
  ConflictException,
  Inject,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { CreateUserDto } from 'src/users/dtos/create-user';
import { UsersService } from 'src/users/users/users.service';
import { LoginDto } from '../dtos/login-dto';
import { UtilsService } from '../utils/utils.service';
import { Cache } from 'cache-manager';
import { VerifyDto } from '../dtos/verify-dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UsersService,
    private readonly jwtService: JwtService,
    private readonly utils: UtilsService,
    @Inject(CACHE_MANAGER) private readonly cacheManager: Cache,
  ) {}

  async signin(dto: LoginDto) {
    const user = await this.userService.findUserByEmail(dto.email);

    if (!this.utils.compareHash(dto.password, user.password)) {
      throw new UnauthorizedException();
    }

    const payload = { email: user.id };

    const accessToken = this.jwtService.sign(payload);

    await this.cacheManager.set(`AUTH_${accessToken}`, user.id, { ttl: 3600 });

    return {
      access_token: accessToken,
    };
  }

  async verify(dto: VerifyDto) {
    const email = (await this.cacheManager.get(`AUTH_${dto.token}`)) as string;

    if (!email) {
      throw new UnauthorizedException();
    }

    return {
      user_email: email,
    };
  }

  async signout(token: string) {
    if (token.split(' ').length != 2) {
      throw new UnauthorizedException();
    }
    

    await this.cacheManager.del(`AUTH_${token.split(' ')[1]}`);
  }

  async signup(dto: CreateUserDto) {
    const user = await this.userService.findUserByEmail(dto.email);

    if (user) {
      throw new ConflictException('user already exists');
    }

    dto.password = this.utils.hashPassword(dto.password);

    await this.userService.create(dto);

    const payload = { email: dto.email };

    const accessToken = this.jwtService.sign(payload);

    await this.cacheManager.set(`AUTH_${accessToken}`, dto.email, {
      ttl: 3600,
    });

    return {
      access_token: accessToken,
    };
  }
}

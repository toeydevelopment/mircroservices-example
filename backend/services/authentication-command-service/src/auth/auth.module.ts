import { CacheModule, Module } from '@nestjs/common';
import type { RedisClientOptions } from 'redis';
import * as redisStore from 'cache-manager-redis-store';

import { AuthService } from './auth/auth.service';
import { AuthController } from './auth/auth.controller';
import { UsersModule } from 'src/users/users.module';
import { JwtModule } from '@nestjs/jwt';
import { UtilsService } from './utils/utils.service';
import { JwtStrategy } from './auth/jwt-strategy';
import { UsersService } from 'src/users/users/users.service';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    ConfigModule.forRoot(),
    UsersModule,
    CacheModule.register<RedisClientOptions<any, any>>({
      url: 'redis://' + process.env.REDIS_HOST,
      store: redisStore,
    }),
    JwtModule.register({
      secret: (() => {
        return process.env.JWT_SECRET;
      })(),
      signOptions: { expiresIn: '60s' },
    }),
  ],
  providers: [AuthService, UtilsService, JwtStrategy],
  controllers: [AuthController],
})
export class AuthModule {}

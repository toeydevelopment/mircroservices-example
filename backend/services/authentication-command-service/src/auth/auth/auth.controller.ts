import { Body, Controller, HttpCode, Post, Req } from '@nestjs/common';
import { CreateUserDto } from 'src/users/dtos/create-user';
import { LoginDto } from '../dtos/login-dto';
import { VerifyDto } from '../dtos/verify-dto';
import { AuthService } from './auth.service';
import { Request } from 'express';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('verify')
  @HttpCode(200)
  verifyToken(@Body() dto: VerifyDto) {
    return this.authService.verify(dto);
  }

  @Post('signin')
  @HttpCode(201)
  signIn(@Body() dto: LoginDto) {
    return this.authService.signin(dto);
  }

  @Post('signup')
  @HttpCode(201)
  signUp(@Body() dto: CreateUserDto) {
    return this.authService.signup(dto);
  }

  @Post('signout')
  @HttpCode(201)
  signOut(@Req() req: Request) {
    return this.authService.signout(req.headers.authorization);
  }
}

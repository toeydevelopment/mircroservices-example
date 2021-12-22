import { IsJWT, IsNotEmpty } from 'class-validator';

export class VerifyDto {
  @IsJWT()
  @IsNotEmpty()
  token: string;
}

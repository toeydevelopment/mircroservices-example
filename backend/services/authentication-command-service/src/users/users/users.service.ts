import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CreateUserDto } from '../dtos/create-user';
import { User, UserDocument } from '../schemas/user';

@Injectable()
export class UsersService {
  constructor(
    @InjectModel(User.name) private readonly userModel: Model<UserDocument>,
  ) {}

  async create(dto: CreateUserDto): Promise<User> {
    const createdUser = new this.userModel({
      _id: dto.email,
      password: dto.password,
      createdAt: new Date(),
    });

    return createdUser.save();
  }


  async findUserByEmail(email: string): Promise<User> {
      return this.userModel.findById(email)
  }
}

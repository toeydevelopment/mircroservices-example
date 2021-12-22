import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, Schema as S } from 'mongoose';

export type UserDocument = User & Document;

@Schema()
export class User {
  @Prop({
    OptionsConstructor: {
      _id: true,
      name: '_id',
    },
  })
  id: string;

  @Prop()
  password: string;

  @Prop()
  createdAt: Date;

  @Prop()
  updatedAt?: Date;

  @Prop()
  deletedAt?: Date;
}

export const UserSchema = new S({
  _id: {
    type: String,
  },
  password: {
    type: String,
  },
  createdAt: {
    type: Date,
  },
  deletedAt: {
    type: Date,
  },
  updatedAt: {
    type: Date,
  },
});

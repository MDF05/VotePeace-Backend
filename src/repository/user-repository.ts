import { User } from "@prisma/client";
import prisma from "../libs/prisma";
import { RegisterDTO } from "../DTO/register-DTO";

class UserRepository {
  async findByNik(nik: string): Promise<User | null> {
    return prisma.user.findUnique({ where: { nik } });
  }

  async createUser(dto: RegisterDTO & { password: string }): Promise<User> {
    return prisma.user.create({ data: dto });
  }
}

export default new UserRepository();

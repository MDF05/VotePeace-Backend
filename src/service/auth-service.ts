import { User } from "@prisma/client";
import { LoginDTO, LoginResponseDTO } from "../DTO/login-DTO";
import { RegisterDTO } from "../DTO/register-DTO";
import userRepository from "../repository/user-repository";

import jwt from "jsonwebtoken";
import dotenv from "dotenv";
import bcrypt from "bcrypt";
dotenv.config();

class AuthService {
  async login(dto: LoginDTO): Promise<LoginResponseDTO> {
    const user: User | null = await userRepository.findByNik(dto.nik);
    if (!user) throw new Error("NIK atau password salah");

    const isPasswordMatch: boolean = await bcrypt.compare(dto.password, user.password);
    if (!isPasswordMatch) throw new Error("NIK atau password salah");

    const { password, ...userWithoutPassword } = user;
    const token: string = jwt.sign(userWithoutPassword, process.env.JWT_SECRET as string, {
      expiresIn: "1d",
    });

    return { user: userWithoutPassword, token };
  }

  async register(dto: RegisterDTO): Promise<Omit<User, "password">> {
    const existUser: User | null = await userRepository.findByNik(dto.nik);
    if (existUser) throw new Error("User dengan NIK ini sudah ada");

    const hashedPassword = await bcrypt.hash(dto.password, 10);

    const user: User = await userRepository.createUser({
      nik: dto.nik,
      name: dto.name,
      password: hashedPassword,
      role: "USER"
    });

    const { password, ...userWithoutPassword } = user;
    return userWithoutPassword;
  }
}


// Trigger restart
export default new AuthService();

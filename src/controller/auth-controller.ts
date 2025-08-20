import { NextFunction, Request, Response } from "express";
import createError from "../utils/create-error";
import { LoginDTO } from "../DTO/login-DTO";
import authService from "../service/auth-service";
import successResponse from "../utils/success-response";
import { RegisterDTO } from "../DTO/register-DTO";

class AuthController {
  async login(req: Request, res: Response, next: NextFunction) {
    try {
      const body: LoginDTO = req.body;
      const loginResponse = await authService.login(body);

      res
        .status(200)
        .json(successResponse("User successfully login", loginResponse));
    } catch (err: unknown) {
      if (err instanceof Error) next(createError(err.message, 400));
      else next(createError("An error occurred", 500));
    }
  }

  async register(req: Request, res: Response, next: NextFunction) {
    try {
      const body: RegisterDTO = req.body;
      const registerResponse = await authService.register(body);

      res
        .status(200)
        .json(
          successResponse("User successfully registered", registerResponse)
        );
    } catch (err) {
      if (err instanceof Error) next(createError(err.message, 400));
      else next(createError("An error occurred", 500));
    }
  }
}

export default new AuthController();

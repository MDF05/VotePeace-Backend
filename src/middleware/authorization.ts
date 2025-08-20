import { NextFunction, Request, Response } from "express";
import createError from "../utils/create-error";

export default function authorization(requiredRole: "ADMIN" | "USER") {
  return (req: Request, res: Response, next: NextFunction) => {
    const user = (req as any).user;
    if (!user)
      throw createError("Unauthorized, user not found in request", 401);

    if (user.role !== requiredRole) {
      throw createError("Forbidden, insufficient role", 403);
    }

    next();
  };
}

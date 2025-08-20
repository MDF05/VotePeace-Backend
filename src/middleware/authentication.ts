import { NextFunction, Request, Response } from "express";
import jwt, { JwtPayload } from "jsonwebtoken";
import createError from "../utils/create-error";

interface DecodedUser extends JwtPayload {
  id: string;
  nik: string;
  name: string;
  role: string;
}

export default function authentication(
  req: Request,
  res: Response,
  next: NextFunction
) {
  try {
    const token = req.headers.authorization?.split("Bearer ")[1];
    if (!token) throw createError("Unauthorized, token missing", 401);

    const decoded = jwt.verify(
      token,
      process.env.JWT_SECRET as string
    ) as DecodedUser;
    if (!decoded) throw createError("Invalid token", 401);

    // inject ke request
    (req as any).user = decoded;
    next();
  } catch (err) {
    next(createError("Unauthorized, invalid token", 401));
  }
}

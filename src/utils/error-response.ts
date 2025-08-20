import { NextFunction, Request, Response } from "express"
import dotenv from "dotenv"
import { CustomError } from "./create-error"
dotenv.config()

export default function errorResponse(err: CustomError, req: Request, res: Response, next: NextFunction) {
    const message = err.message
    const stack = err.stack
    const name = err.name
    const status = err.status
    const succes = err.succes

    return res.status(status).json({
        succes,
        author: "MUHAMMAD DAVA FAHREZA",
        aplication: "VotePeace",
        version: process.env.version,
        message,
        date: new Date(),
        status,
        content: [
            {
                name,
                protocol: req.protocol,
                host: req.get("host"),
                endPoint: req.originalUrl,
                method: req.method,
                userAgent: req.headers["user-agent"],
                ip: req.ip,
                ips: req.ips,
                stack,
            }]


    })
}
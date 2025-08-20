import { successResponseTypes } from "./types/succes-response-types";

export default function successResponse(message: string, content: object, status?: number): successResponseTypes {
    return {
        succes: true,
        author: "MUHAMMAD DAVA FAHREZA",
        aplication: "VotePeace",
        version: process.env.version,
        message,
        date: new Date(),
        status: status || 200,
        content

    }
}
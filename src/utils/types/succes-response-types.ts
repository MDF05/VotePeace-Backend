export interface successResponseTypes {
    succes: boolean,
    author: string,
    aplication: string,
    version: string | undefined,
    message: string,
    date: Date,
    status: number,
    content: object
}
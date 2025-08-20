export interface RegisterDTO {
  nik: string;
  name: string;
  password: string;
  role?: "USER" | "ADMIN";
}

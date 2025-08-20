export interface UserDTO {
  id: string;
  nik: string;
  name: string;
  role: "USER" | "ADMIN";
  createdAt: Date;
  updatedAt: Date;
}

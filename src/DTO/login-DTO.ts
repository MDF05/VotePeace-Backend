export interface LoginDTO {
  nik: string;
  password: string;
}

export interface LoginResponseDTO {
  user: {
    id: string;
    nik: string;
    name: string;
    role: "USER" | "ADMIN";
    createdAt: Date;
    updatedAt: Date;
  };
  token: string;
}

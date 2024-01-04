CREATE TABLE "produtos" (
    "id" serial PRIMARY KEY NOT NULL,
    "nome" VARCHAR(255) NOT NULL,
    "descricao" VARCHAR(255) NOT NULL,
    "preco" INTEGER NOT NULL,
    "criado_em" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
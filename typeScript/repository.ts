import * as fs from "fs/promises";
export class NoteRepository {
  constructor() {
    fs.readFile("./db.txt", "utf8").catch((err) => {
      if (err) {
        fs.writeFile("./db.txt", "", "utf8");
      }
    });
  }

  async save(data: string): Promise<void> {
    await fs.appendFile("./db.txt", `${data}\n`, "utf8");
  }

  async findAll(): Promise<string> {
    const file = await fs.readFile("./db.txt");
    return file.toString();
  }
}

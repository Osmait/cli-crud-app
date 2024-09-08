import * as fs from "fs/promises";
import { Note } from "./note.model";
export class NoteRepository {
  private fileLocaltion = "./";
  private fileName = "db.json";
  private fileReferes = `${this.fileLocaltion}${this.fileName}`;
  constructor() {
    this.initializeFile();
  }
  private async initializeFile() {
    try {
      await fs.readFile(this.fileReferes, "utf8");
    } catch (err) {
      if (err.code === "ENOENT") {
        await fs.writeFile(this.fileReferes, "[]", "utf8");
        console.log("Archivo creado con contenido inicial vacío.");
      } else {
        console.error("Error inesperado al leer el archivo:", err);
        throw err;
      }
    }
  }

  async save(data: Note): Promise<void> {
    const dataDb = await fs.readFile(this.fileReferes, "utf8");
    const dataTosave: Note[] = dataDb.trim() ? JSON.parse(dataDb) : [];
    dataTosave.push(data);
    const jsonString = JSON.stringify(dataTosave, null, 2);

    await fs.writeFile(this.fileReferes, `${jsonString}\n`, "utf8");
  }

  async findAll(): Promise<Note[]> {
    const file = await fs.readFile(this.fileReferes, "utf8");
    return JSON.parse(file);
  }
}

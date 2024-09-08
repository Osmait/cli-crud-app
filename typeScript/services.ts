import { Note } from "./note.model";
import { NoteRepository } from "./repository";

export class NoteServices {
  private noteRepository: NoteRepository;

  constructor(noteRepository: NoteRepository) {
    this.noteRepository = noteRepository;
  }
  async save(note: string, name: string) {
    const noteToSave: Note = {
      id: Math.random().toString(),
      name: name,
      description: note,
      createdAt: Date.now().toString(),
      updateAt: Date.now().toString(),
    };
    await this.noteRepository.save(noteToSave);
  }

  async findAll(): Promise<Note[]> {
    return await this.noteRepository.findAll();
  }
}

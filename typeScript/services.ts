import { NoteRepository } from "./repository";

export class NoteServices {
  private noteRepository: NoteRepository;

  constructor(noteRepository: NoteRepository) {
    this.noteRepository = noteRepository;
  }
  async save(note: string) {
    await this.noteRepository.save(note);
  }

  async findAll() {
    await this.findAll();
  }
}

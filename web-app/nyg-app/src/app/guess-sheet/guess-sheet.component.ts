import { Component, Input, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { DragDropModule } from '@angular/cdk/drag-drop';
@Component({
  selector: 'app-guess-sheet',
  imports: [MatCardModule, DragDropModule],
  templateUrl: './guess-sheet.component.html',
  styleUrl: './touchup/guess-sheet.component.scss',
  encapsulation: ViewEncapsulation.Emulated
})
export class GuessSheetComponent {
  @Input() GuessSheetValue: WritableSignal<string> = signal("")
}

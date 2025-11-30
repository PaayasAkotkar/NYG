import { Component, Input, signal, WritableSignal } from '@angular/core';
import { ListComponent } from "../list/list.component";
import { OCSFieldFormat } from '../../helpers/patterns';

@Component({
  selector: 'app-frame',
  imports: [ListComponent],
  templateUrl: './frame.component.html',
  styleUrl: './frame.component.scss'
})
export class OCSFrameComponent {
  @Input() OCSlist: WritableSignal<OCSFieldFormat[]> = signal([

  ])
  @Input() OCSproceed: WritableSignal<boolean> = signal(false)
  Header = 'FRAMES'
}

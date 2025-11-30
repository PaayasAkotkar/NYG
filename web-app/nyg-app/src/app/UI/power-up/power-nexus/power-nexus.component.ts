import { Component, Input, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../../lobby/resources/lobbyResource';
import { RoomService } from '../../../roomService/room.service';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { MatCardModule } from '@angular/material/card';

@Component({
  selector: 'app-power-nexus',
  imports: [DragDropModule, MatCardModule],
  templateUrl: './power-nexus.component.html',
  styleUrls: ['./touchup/nexus.scss', './touchup/animation.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class PowerNexusComponent {
  @Input() Nexus: WritableSignal<string> = signal("")
  constructor() {
  }
}

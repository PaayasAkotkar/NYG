import { Directive, Output, EventEmitter, HostListener } from '@angular/core';
import { SetupMenuService } from '../setupMenu/setup-menu.service';

@Directive({
  selector: '[appCloseMenu]'
})
export class CloseMenuDirective {

  @Output() NYGcloseMenu = new EventEmitter<KeyboardEvent>();
  @HostListener('document:keydown', ['$event'])
  handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Control') {
      this.NYGcloseMenu.emit(event);

    }
  }
}

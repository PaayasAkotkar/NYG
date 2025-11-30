import { Component } from '@angular/core';

@Component({
  selector: 'app-material',
  imports: [],
  templateUrl: './material.component.html',
  styleUrls: ['./touchup/material.scss']
})
export class MaterialComponent {
  // todo: remove events that are chosen
  // important: because 0 index is considered nor even not odd
  NYGtitles = [
  ].sort()
  odd = this.NYGtitles.filter((e, i) => i % 2 == 0 && e != '')
  even = this.NYGtitles.filter((e, i) => i % 2 != 0 && e != ' ')
}

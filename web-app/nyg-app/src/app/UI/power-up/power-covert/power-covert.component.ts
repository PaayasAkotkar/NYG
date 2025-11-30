import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-power-covert',
  imports: [],
  templateUrl: './power-covert.component.html',
  styleUrls: ['./touchup/covert.scss', './touchup/animations.scss', './touchup/color.scss']
})
export class PowerCovertComponent implements OnInit {

  constructor() { }

  particles: number[] = []

  ngOnInit(): void {
    for (let i = 0; i < 10; i++) {
      this.particles.push(i)
    }
  }
}

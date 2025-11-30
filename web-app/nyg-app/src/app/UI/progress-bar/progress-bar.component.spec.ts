import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProgressBarComponent } from './progress-bar.component';

describe('ProgressBarComponent', () => {
  let component: ProgressBarComponent;
  let fixture: ComponentFixture<ProgressBarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProgressBarComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(ProgressBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should set default img incase no img send', () => {
    const src = component.OCSimgSrc = ''
    fixture.detectChanges()

    expect(src).toBe(component.defaultPic)

    const dummy = fixture.nativeElement as HTMLElement
    const dummyImg = dummy.querySelector('img') as HTMLImageElement
    expect(dummyImg).toBeTruthy()
    expect(dummyImg.src).toBe(src)

  })
});

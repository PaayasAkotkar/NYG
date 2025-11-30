import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ClashLobbyComponent } from './clash-lobby.component';
import { Location } from '@angular/common';

describe('ClashLobbyComponent', () => {
  let component: ClashLobbyComponent;
  let fixture: ComponentFixture<ClashLobbyComponent>;

  beforeEach(async () => {

    await TestBed.configureTestingModule({
      imports: [ClashLobbyComponent],
      providers: [Location,]
    })
      .compileComponents();

    fixture = TestBed.createComponent(ClashLobbyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  let locationSpy = TestBed.inject(Location) as jasmine.SpyObj<Location>

  it('should exaclty set two keys from the predefinedKeys', () => {
    fixture.detectChanges()
    const predefinedKeys = ['nexus', 'covert', 'rewind', 'bet', 'freeze'];
    const selectedKeys = predefinedKeys.slice(0, 2); // Select any two keys

    locationSpy.getState.and.returnValue({ Keys: selectedKeys });
    fixture.detectChanges(); // triggers ngOnInit

  })
  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

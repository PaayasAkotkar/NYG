import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { ApplicationModule } from '@angular/core';
import { ClickScrollPlugin, OverlayScrollbars } from 'overlayscrollbars';

OverlayScrollbars.plugin(ClickScrollPlugin)

bootstrapApplication(AppComponent, appConfig

)
  .catch((err) => console.error(err));

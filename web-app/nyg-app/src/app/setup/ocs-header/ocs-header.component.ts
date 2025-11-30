import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ocs-header',
  imports: [],
  templateUrl: './ocs-header.component.html',
  styleUrls: ['./touchup/ocs-header.scss']
})
export class OcsHeaderComponent {
  @Input() OCStitle: string = "STAY UPDATED"
  @Input() OCSsubTitles: any[] | string = ['YOUTUBE', 'X', 'INSTA']
  @Input() isHeader: boolean = false
  @Input() isLink: boolean = false
  @Input() OCShref: string = ""
  @Input() NYGerrorMessage: string
  @Input() NYGisError: boolean = true
  @Input() NYGsetToolTip: boolean = false
  @Input() NYGtoolTip: string = ""
}

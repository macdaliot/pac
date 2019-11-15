import { Component, OnInit, Input } from '@angular/core';
import { faInfoCircle, faExclamationTriangle, faExclamationCircle, faCheckCircle } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-alert-card',
  templateUrl: './alert-card.component.html',
  styleUrls: ['./alert-card.component.scss']
})
export class AlertCardComponent implements OnInit {
  faInfoCircle = faInfoCircle;
  faExclamationTriangle = faExclamationTriangle;
  faExclamationCircle = faExclamationCircle;
  faCheckCircle = faCheckCircle;

  alertIcon;

  private _alertLevel = 'info';

  get alertLevel() {
  	return this._alertLevel;
  }
  @Input()
  set alertLevel(value: string) {
  	this._alertLevel = value;

  	if (value === 'info') {
  		this.alertIcon = this.faInfoCircle;
  	} else if (value === 'warn') {
  		this.alertIcon = this.faExclamationTriangle;
  	} else if (value === 'error') {
  		this.alertIcon = this.faExclamationCircle;
  	} else if (value === 'success') {
  		this.alertIcon = this.faCheckCircle;
  	}

  }

  constructor() {
  	this.alertIcon = this.faCheckCircle;
  }

  ngOnInit() {
  }

}

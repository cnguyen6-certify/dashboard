// Copyright 2017 The Kubernetes Dashboard Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {StateParams} from 'common/resource/resourcedetail';
import {stateName} from 'replicationcontroller/detail/state';

/**
 * Controller for the replication controller card menu
 *
 * @final
 */
export default class ReplicationControllerCardMenuController {
  /**
   * @param {!ui.router.$state} $state
   * @param {!./../../common/scaling/service.ScaleService} kdScaleService
   * @ngInject
   */
  constructor($state, kdScaleService) {
    /**
     * Initialized from the scope.
     * @export {!backendApi.ReplicationController}
     */
    this.replicationController;

    /** @private {!ui.router.$state} */
    this.state_ = $state;

    /** @private {!./../../common/scaling/service.ScaleService} */
    this.kdScaleService_ = kdScaleService;
  }

  /**
   * @param {!function(!MouseEvent)} $mdOpenMenu
   * @param {!MouseEvent} $event
   * @export
   */
  openMenu($mdOpenMenu, $event) {
    $mdOpenMenu($event);
  }

  /**
   * @export
   */
  viewDetails() {
    this.state_.go(
        stateName,
        new StateParams(
            this.replicationController.objectMeta.namespace,
            this.replicationController.objectMeta.name));
  }

  /**
   * @export
   */
  showScaleDialog() {
    this.kdScaleService_.showScaleDialog(
        this.replicationController.objectMeta.namespace, this.replicationController.objectMeta.name,
        this.replicationController.pods.current, this.replicationController.pods.desired,
        this.replicationController.typeMeta.kind);
  }
}

/**
 * @return {!angular.Component}
 */
export const replicationControllerCardMenuComponent = {
  bindings: {
    'replicationController': '=',
  },
  controller: ReplicationControllerCardMenuController,
  templateUrl: 'replicationcontroller/list/cardmenu.html',
};

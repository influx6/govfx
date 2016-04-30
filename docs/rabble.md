# Rabble
  Contains messy information which are internal used and not for external purpose.


init initialize the timer internal structures.
-------------------------------------------------------------------------

 tween resumption: startTime = _resumeTime - Math.abs(shift) - _progressTime;
 tween prevtime:   startTime + _progressTime - delay;
 tween Dimensions:  time = duration + delay | repeatTime = time * (repeat + 1);
-------------------------------------------------------------------------
StartTime
   set start time to passed time or to the current moment
   var startTime = (time == null) ? performance.now() : time;
   calculate bounds
   - negativeShift is negative delay in options hash
   - shift time is shift of the parent
   p.startTime = startTime + p.delay + this._negativeShift + shiftTime;
   p.endTime   = p.startTime + p.repeatTime - p.delay;
   set play time to the startTime
   if playback controls are used - use _resumeTime as play time,
   else use shifted startTime -- shift is needed for timelines append chains
   this._playTime = ( this._resumeTime != null )
     ? this._resumeTime : startTime + shiftTime;
-------------------------------------------------------------------------
UpdateTime
  var startPoint = p.startTime - p.delay;

  // if speed param was defined - calculate
    // new time regarding speed
    if ( p.speed && this._playTime ) {
      // play point + ( speed * delta )
      time = this._playTime + ( p.speed * ( time - this._playTime ) );
    }

  if in active area and not ended - save progress time
  for pause/play purposes.
  if ( time > startPoint && time < p.endTime ) {
    this._progressTime = time - startPoint;
  }
  else if not started or ended set progress time to 0
  else if ( time <= startPoint  ) { this._progressTime = 0; }
  else if ( time >= p.endTime ) {
    set progress time to repeat time + tiny cofficient
    to make it extend further than the end time
    this._progressTime = p.repeatTime + .00000000001;
  }

  reverse time if _props.isReversed is set
  if ( p.isReversed ) { time = p.endTime - this._progressTime; }
   We need to know what direction we are heading to,
   so if we don't have the previous update value - this is very first
   update, - skip it entirely and wait for the next value

   handle onProgress callback
     if  ( time >= startPoint && time <= p.endTime ) {
       this._progress( (time - startPoint) / p.repeatTime, time );
     }

       if time is inside the active area of the tween.
       active area is the area from start time to end time,
       with all the repeat and delays in it

     if ((time >= p.startTime) && (time <= p.endTime)) {
       this._updateInActiveArea( time );
     } else { (this._isInActiveArea) && this._updateInInactiveArea( time ); }

     this._prevTime = time;
     return (time >= p.endTime) || (time <= startPoint);

-------------------------------------------------------------------------

UpdateInAcive
      var props         = this._props,
      delayDuration = props.delay + props.duration,
       startPoint    = props.startTime - props.delay,
       elapsed       = (time - props.startTime + props.delay) % delayDuration,
       TCount        = Math.round( (props.endTime - props.startTime + props.delay) / delayDuration ),
       T             = this._getPeriod(time),
       TValue        = this._delayT,
       prevT         = this._getPeriod(this._prevTime),
       TPrevValue    = this._delayT;

   if time is inside the duration area of the tween
  if ( startPoint + elapsed >= props.startTime ) {
    this._isInActiveArea = true; this._isRepeatCompleted = false;
    this._isRepeatStart = false; this._isStarted = false;
    // active zone or larger then end
    var elapsed2 = ( time - props.startTime) % delayDuration,
        proc = elapsed2 / props.duration;
    // |=====|=====|=====| >>>
    //      ^1^2
    var isOnEdge = (T > 0) && (prevT < T);
    // |=====|=====|=====| <<<
    //      ^2^1
    var isOnReverseEdge = (prevT > T);

  _getPeriod ( time ) {
     var p       = this._props,
         TTime   = p.delay + p.duration,
         dTime   = p.delay + time - p.startTime,
         T       = dTime / TTime,
         // if time if equal to endTime we need to set the elapsed
         // time to 0 to fix the occasional precision js bug, which
         // causes 0 to be something like 1e-12
         elapsed = ( time < p.endTime ) ? dTime % TTime : 0;
     // If the latest period, round the result, otherwise floor it.
     // Basically we always can floor the result, but because of js
     // precision issues, sometimes the result is 2.99999998 which
     // will result in 2 instead of 3 after the floor operation.
     T = ( time >= p.endTime ) ? Math.round(T) : Math.floor(T);
     // if time is larger then the end time
     if ( time > p.endTime ) {
       // set equal to the periods count
       T = Math.round( (p.endTime - p.startTime + p.delay) / TTime );
     // if in delay gap, set _delayT to current
     // period number and return "delay"
     } else if ( elapsed > 0 && elapsed < p.delay ) {
       this._delayT = T; T = 'delay';
     }
     // if the end of period and there is a delay
     return T;
   }
-------------------------------------------------------------------------

*/

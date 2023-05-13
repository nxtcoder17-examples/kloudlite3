import { useRef} from "react";
import PropTypes from "prop-types";
import { BounceIt } from "../bounce-it";
import classnames from 'classnames';
import { useCheckbox } from "react-aria";
import { useToggleState } from "react-stately";


export const Checkbox = ({ label, error, ...props})=>{
  const ref = useRef();
  const state = useToggleState(props);
  const disabled = props.disabled;
  const {inputProps} = useCheckbox(props, state, ref);
  return <BounceIt disable={props.disabled}>
    <div    
      className={"flex gap-2 select-none items-center relative"} 
      onClick={()=>{
        if(ref.current){
          ref.current.click();
        }
      }}
    >
    <div className="relative h-5">
      <input 
        {...inputProps}
        className={classnames(
          "appearance-none w-5 h-5 border rounded",
          "disabled:cursor-not-allowed disabled:opacity-50 disabled:bg-surface-default disabled:border-border-disabled",
          {
            "border-border-default  checked:border-border-primary indeterminate:border-border-primary checked:bg-surface-primary-default indeterminate:bg-surface-primary-default bg-surface-default":!error && !disabled,
            "border-border-danger checked:border-border-danger indeterminate:border-border-danger checked:bg-surface-danger-default indeterminate:bg-surface-danger-default bg-surface-danger-subdued":error,
          },
          "ring-offset-1",
          "transition-all",
          "outline-none",
          "focus-visible:ring-2 focus:ring-border-focus",
        )}
        ref={ref}
      />
      {
        props.isIndeterminate && <div className="absolute top-2/4 left-2/4 -translate-y-2/4 -translate-x-2/4">
            <svg width="17" height="16" viewBox="0 0 17 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M2.5 8H14.5" stroke="#F9FAFB" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </div>
      }
      {
        (!props.isIndeterminate && state.isSelected) && <div className="absolute top-2/4 left-2/4 -translate-y-2/4 -translate-x-2/4">
          <svg width="17" height="16" viewBox="0 0 17 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M14.5 4.00019L6.5 11.9998L2.5 8.00019" stroke="#F9FAFB" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
      }
    </div>
    <label className={classnames({
      "text-text-default":!disabled,
      "text-text-disabled":disabled,
    }, "bodyMd-medium")}>{label}</label>
  </div>
  </BounceIt>
}

Checkbox.propTypes = {
  value: PropTypes.oneOf(["true", "false", "indeterminate"]),
  onChange: PropTypes.func,
  label: PropTypes.string.isRequired,
  disabled: PropTypes.bool,
  error: PropTypes.bool,
  isIndeterminate: PropTypes.bool,
}

Checkbox.defaultProps = {
  onChange: ()=>{},
  disabled: false,
  error: false,
  isIndeterminate: false,
}
